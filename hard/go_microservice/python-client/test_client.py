import unittest
from unittest.mock import patch, MagicMock, PropertyMock
import sys
import os

sys.path.insert(0, os.path.dirname(__file__))

try:
    import integration_pb2
    import integration_pb2_grpc
    from client import IntegrationClient
except ImportError:
    integration_pb2 = MagicMock()
    integration_pb2_grpc = MagicMock()
    from client import IntegrationClient


class TestIntegrationClientHTTP(unittest.TestCase):
    @patch('client.requests.post')
    def test_integrate_http_sin(self, mock_post):
        mock_response = MagicMock()
        mock_response.json.return_value = {
            "result": 2.0,
            "error_estimate": 1e-10,
            "partitions": 100000
        }
        mock_post.return_value = mock_response

        client = IntegrationClient(http_host="localhost:18080")
        result = client.integrate_http("sin", 0, 3.14159, 100000)

        self.assertEqual(result["result"], 2.0)
        self.assertEqual(result["partitions"], 100000)
        mock_post.assert_called_once()

    @patch('client.requests.post')
    def test_integrate_http_x_squared(self, mock_post):
        mock_response = MagicMock()
        mock_response.json.return_value = {
            "result": 0.333333,
            "error_estimate": 1e-11,
            "partitions": 50000
        }
        mock_post.return_value = mock_response

        client = IntegrationClient()
        result = client.integrate_http("x^2", 0, 1, 50000)

        self.assertAlmostEqual(result["result"], 0.333333, places=5)
        mock_post.assert_called_once()

    @patch('client.requests.post')
    def test_integrate_http_exp(self, mock_post):
        mock_response = MagicMock()
        mock_response.json.return_value = {
            "result": 1.718281,
            "error_estimate": 1e-9,
            "partitions": 10000
        }
        mock_post.return_value = mock_response

        client = IntegrationClient()
        result = client.integrate_http("exp", 0, 1, 10000)

        self.assertAlmostEqual(result["result"], 1.718281, places=4)

    @patch('client.requests.post')
    def test_integrate_http_cos(self, mock_post):
        mock_response = MagicMock()
        mock_response.json.return_value = {
            "result": 1.0,
            "error_estimate": 1e-10,
            "partitions": 50000
        }
        mock_post.return_value = mock_response

        client = IntegrationClient()
        result = client.integrate_http("cos", 0, 1.5708, 50000)

        self.assertAlmostEqual(result["result"], 1.0, places=3)

    @patch('client.requests.post')
    def test_integrate_http_raises_on_error(self, mock_post):
        mock_response = MagicMock()
        mock_response.raise_for_status.side_effect = Exception("HTTP Error")
        mock_post.return_value = mock_response

        client = IntegrationClient()
        with self.assertRaises(Exception):
            client.integrate_http("sin", 0, 1, 1000)


class TestIntegrationClientGRPC(unittest.TestCase):
    def _create_mock_client(self):
        mock_channel = MagicMock()
        mock_stub = MagicMock()
        
        def mock_unary_unary(*args, **kwargs):
            return lambda request: self._create_response(request)
        
        mock_channel.unary_unary = MagicMock(side_effect=mock_unary_unary)
        return mock_channel, mock_stub

    def _create_response(self, request):
        response = MagicMock()
        response.result = 2.0
        response.error_estimate = 1e-10
        response.partitions_used = request.partitions if hasattr(request, 'partitions') else 100000
        return response

    @patch.object(integration_pb2_grpc, 'IntegratorStub')
    @patch('client.grpc.insecure_channel')
    def test_integrate_grpc_sin(self, mock_channel, mock_stub_class):
        mock_stub = MagicMock()
        mock_response = MagicMock()
        mock_response.result = 2.0
        mock_response.error_estimate = 1e-10
        mock_response.partitions_used = 100000
        mock_stub.Integrate.return_value = mock_response
        mock_stub_class.return_value = mock_stub
        mock_channel.return_value = MagicMock()

        client = IntegrationClient()
        result = client.integrate_grpc("sin", 0, 3.14159, 100000)

        self.assertEqual(result["result"], 2.0)
        self.assertEqual(result["partitions"], 100000)
        mock_stub.Integrate.assert_called_once()

    @patch.object(integration_pb2_grpc, 'IntegratorStub')
    @patch('client.grpc.insecure_channel')
    def test_integrate_grpc_x_squared(self, mock_channel, mock_stub_class):
        mock_stub = MagicMock()
        mock_response = MagicMock()
        mock_response.result = 0.333333
        mock_response.error_estimate = 1e-11
        mock_response.partitions_used = 50000
        mock_stub.Integrate.return_value = mock_response
        mock_stub_class.return_value = mock_stub
        mock_channel.return_value = MagicMock()

        client = IntegrationClient()
        result = client.integrate_grpc("x^2", 0, 1, 50000)

        self.assertAlmostEqual(result["result"], 0.333333, places=5)

    @patch.object(integration_pb2_grpc, 'IntegratorStub')
    @patch('client.grpc.insecure_channel')
    def test_integrate_grpc_exp(self, mock_channel, mock_stub_class):
        mock_stub = MagicMock()
        mock_response = MagicMock()
        mock_response.result = 1.718281
        mock_response.error_estimate = 1e-9
        mock_response.partitions_used = 10000
        mock_stub.Integrate.return_value = mock_response
        mock_stub_class.return_value = mock_stub
        mock_channel.return_value = MagicMock()

        client = IntegrationClient()
        result = client.integrate_grpc("exp", 0, 1, 10000)

        self.assertAlmostEqual(result["result"], 1.718281, places=4)

    @patch.object(integration_pb2_grpc, 'IntegratorStub')
    @patch('client.grpc.insecure_channel')
    def test_integrate_grpc_stream(self, mock_channel, mock_stub_class):
        mock_stub = MagicMock()
        mock_responses = [
            MagicMock(result=1.0, error_estimate=1e-5, partitions_used=1000),
            MagicMock(result=1.5, error_estimate=1e-6, partitions_used=2000),
            MagicMock(result=1.718, error_estimate=1e-7, partitions_used=3000),
        ]
        mock_stub.IntegrateStream.return_value = iter(mock_responses)
        mock_stub_class.return_value = mock_stub
        mock_channel.return_value = MagicMock()

        client = IntegrationClient()
        results = list(client.integrate_grpc_stream("exp", 0, 1, 3000))

        self.assertEqual(len(results), 3)
        self.assertEqual(results[0]["result"], 1.0)
        self.assertEqual(results[2]["result"], 1.718)

    @patch('client.grpc.insecure_channel')
    def test_close_channel(self, mock_channel):
        mock_channel_instance = MagicMock()
        mock_channel.return_value = mock_channel_instance

        client = IntegrationClient()
        client.close()

        mock_channel_instance.close.assert_called_once()


class TestIntegrationClientEnvVars(unittest.TestCase):
    @patch.dict(os.environ, {"HTTP_HOST": "custom:9999", "GRPC_HOST": "custom:55551"})
    @patch('client.requests.post')
    @patch('client.grpc.insecure_channel')
    def test_env_vars_override_defaults(self, mock_channel, mock_post):
        client = IntegrationClient()

        self.assertEqual(client.http_host, "custom:9999")
        mock_channel.assert_called_once_with("custom:55551")


if __name__ == '__main__':
    unittest.main()
