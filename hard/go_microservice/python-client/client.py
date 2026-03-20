import requests
import grpc
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), 'proto'))

try:
    import integration_pb2
    import integration_pb2_grpc
except ImportError:
    print("Proto files not generated. Run: python -m grpc_tools.protoc -I../go-service/proto --python_out=. --grpc_python_out=. ../go-service/proto/integration.proto")
    sys.exit(1)


class IntegrationClient:
    def __init__(self, http_host=None, grpc_host=None):
        http_host = http_host or os.environ.get("HTTP_HOST", "localhost:18080")
        grpc_host = grpc_host or os.environ.get("GRPC_HOST", "localhost:50051")
        self.http_host = http_host
        self.grpc_channel = grpc.insecure_channel(grpc_host)
        self.grpc_stub = integration_pb2_grpc.IntegratorStub(self.grpc_channel)

    def integrate_http(self, func: str, lower: float, upper: float, partitions: int = 10000):
        url = f"http://{self.http_host}/integrate"
        params = {
            "func": func,
            "a": lower,
            "b": upper,
            "n": partitions
        }
        response = requests.post(url, params=params)
        response.raise_for_status()
        return response.json()

    def integrate_grpc(self, func: str, lower: float, upper: float, partitions: int = 10000):
        request = integration_pb2.IntegrationRequest(
            function=func,
            lower_bound=lower,
            upper_bound=upper,
            partitions=partitions
        )
        response = self.grpc_stub.Integrate(request)
        return {
            "result": response.result,
            "error_estimate": response.error_estimate,
            "partitions": response.partitions_used
        }

    def integrate_grpc_stream(self, func: str, lower: float, upper: float, partitions: int = 10000):
        request = integration_pb2.IntegrationRequest(
            function=func,
            lower_bound=lower,
            upper_bound=upper,
            partitions=partitions
        )
        for response in self.grpc_stub.IntegrateStream(request):
            yield {
                "result": response.result,
                "error_estimate": response.error_estimate,
                "partitions": response.partitions_used
            }

    def close(self):
        self.grpc_channel.close()


def main():
    client = IntegrationClient()

    print("=== HTTP Integration ===")
    result = client.integrate_http("sin", 0, 3.14159, 50000)
    print(f"∫sin(x)dx from 0 to π = {result['result']:.10f}")
    print(f"Error estimate: {result['error_estimate']:.2e}")
    print(f"Partitions: {result['partitions']}")
    print()

    print("=== gRPC Integration ===")
    result = client.integrate_grpc("x^2", 0, 1, 50000)
    print(f"∫x²dx from 0 to 1 = {result['result']:.10f}")
    print(f"Error estimate: {result['error_estimate']:.2e}")
    print()

    print("=== gRPC Streaming Integration ===")
    for i, partial in enumerate(client.integrate_grpc_stream("exp", 0, 1, 10000)):
        if i < 3 or i >= 9:
            print(f"Progress {i+1}: result={partial['result']:.6f}")
        elif i == 3:
            print("...")

    client.close()


if __name__ == "__main__":
    main()
