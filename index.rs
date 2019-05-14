use std::ops::Add;

fn add<T: Add<Output=T>>(a: T, b: T) -> T {
    a + b
}

fn main() {
    let sum = add(1, 2);
    println!("{}", sum)
}