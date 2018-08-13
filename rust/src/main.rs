
extern crate image;

use std::fs::File;
use std::io::prelude::*;
use std::{env, process};

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        println!("Usage {} <in_img> <out_img>", args[0]);
        process::exit(-1);
    }

    println!("In file: {}", args[1]);
    println!("Out file: {}", args[2]);

    let mut file = File::open(&args[1])?;

    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    println!("{}", contents);

    let mut out_file = File::create(&args[2])?;
    out_file.write_all(&contents.into_bytes())?;

    Ok(())
}
