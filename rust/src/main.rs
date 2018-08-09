
extern crate image;

use std::{env, process};

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        println!("Usage {} <in_img> <out_img>", args[0]);
        process::exit(-1);
    }

    println!("In file: {}", args[1]);
    println!("Out file: {}", args[2]);
}
