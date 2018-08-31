
extern crate image;

use std::{env, process};
use std::fs::File;
use image::{png, GenericImage, ColorType};

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        println!("Usage {} <in_img> <out_img>", args[0]);
        process::exit(-1);
    }

    let img = image::open(&args[1]).unwrap();

    let buf = File::create(&args[2])?;
    let out_img = png::PNGEncoder::new(buf);
    let (img_width, img_height) = img.dimensions();
    out_img.encode(&(img.raw_pixels()), img_width, img_height, ColorType::RGB(8))?;

    Ok(())
}
