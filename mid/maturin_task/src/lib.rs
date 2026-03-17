use pyo3::prelude::*;

fn sum_as_string_impl(a: usize, b: usize) -> String {
    (a + b).to_string()
}

fn add_impl(a: i64, b: i64) -> i64 {
    a + b
}

fn greet_impl(name: &str) -> String {
    format!("hello, {name}")
}

#[pyfunction]
fn sum_as_string(a: usize, b: usize) -> PyResult<String> {
    Ok(sum_as_string_impl(a, b))
}

#[pyfunction]
fn add(a: i64, b: i64) -> i64 {
    add_impl(a, b)
}

#[pyfunction]
fn greet(name: &str) -> String {
    greet_impl(name)
}

#[pymodule]
fn fastmath(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(sum_as_string, m)?)?;
    m.add_function(wrap_pyfunction!(add, m)?)?;
    m.add_function(wrap_pyfunction!(greet, m)?)?;
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sum_as_string_impl() {
        assert_eq!(sum_as_string_impl(2, 3), "5");
    }

    #[test]
    fn test_add_impl() {
        assert_eq!(add_impl(10, 32), 42);
        assert_eq!(add_impl(-1, 1), 0);
    }

    #[test]
    fn test_greet_impl() {
        assert_eq!(greet_impl("Rust"), "hello, Rust");
    }
}
