
extern crate image;

use std::{env, process};
use image::GenericImage;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        println!("Usage {} <in_img> <out_img>", args[0]);
        process::exit(-1);
    }

    let img = image::open(&args[1]).unwrap();

    println!("Image is {:?}", img.dimensions());

    /*
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    println!("{}", contents);

    let mut out_file = File::create(&args[2])?;
    out_file.write_all(&contents.into_bytes())?;
    */
}
