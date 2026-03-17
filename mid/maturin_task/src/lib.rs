use pyo3::prelude::*;

#[pymodule]
mod fastmath {
    use pyo3::prelude::*;

    #[pyfunction]
    fn sum_as_string(a: usize, b: usize) -> PyResult<String> {
        Ok((a + b).to_string())
    }

    #[pyfunction]
    fn add(a: i64, b: i64) -> i64 {
        a + b
    }

    #[pyfunction]
    fn greet(name: &str) -> String {
        format!("hello, {name}")
    }

    #[pymodule_init]
    fn init(m: &Bound<'_, PyModule>) -> PyResult<()> {
        m.add_function(wrap_pyfunction!(sum_as_string, m)?)?;
        m.add_function(wrap_pyfunction!(add, m)?)?;
        m.add_function(wrap_pyfunction!(greet, m)?)?;
        Ok(())
    }
}
