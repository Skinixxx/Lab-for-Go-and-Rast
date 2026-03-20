use pyo3::prelude::*;

pub fn add_impl(a: i32, b: i32) -> i32 {
    a + b
}

pub fn multiply_impl(a: i32, b: i32) -> i32 {
    a * b
}

pub fn greet_impl(name: &str) -> String {
    format!("Hello, {}!", name)
}

#[pyfunction]
fn add(a: i32, b: i32) -> i32 {
    add_impl(a, b)
}

#[pyfunction]
fn multiply(a: i32, b: i32) -> i32 {
    multiply_impl(a, b)
}

#[pyfunction]
fn greet(name: &str) -> String {
    greet_impl(name)
}

#[pymodule]
fn fastmath(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(add, m)?)?;
    m.add_function(wrap_pyfunction!(multiply, m)?)?;
    m.add_function(wrap_pyfunction!(greet, m)?)?;
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add_impl(2, 3), 5);
        assert_eq!(add_impl(-1, 1), 0);
        assert_eq!(add_impl(0, 0), 0);
        assert_eq!(add_impl(100, 200), 300);
    }

    #[test]
    fn test_multiply() {
        assert_eq!(multiply_impl(2, 3), 6);
        assert_eq!(multiply_impl(-2, 3), -6);
        assert_eq!(multiply_impl(0, 100), 0);
        assert_eq!(multiply_impl(5, 5), 25);
    }

    #[test]
    fn test_greet() {
        assert_eq!(greet_impl("Rust"), "Hello, Rust!");
        assert_eq!(greet_impl("World"), "Hello, World!");
        assert_eq!(greet_impl(""), "Hello, !");
    }
}
