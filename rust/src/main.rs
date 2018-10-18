
extern crate image;
extern crate time;

use std::{env, process};
use std::fs::File;
use image::{png, GenericImage, ColorType, FilterType};
use time;
//use image::{jpeg, png, GenericImage, ImageDecoder, ColorType};

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        println!("Usage {} <in_img> <out_img>", args[0]);
        process::exit(-1);
    }

    let img = image::open(&args[1]).unwrap();
    let (img_width, img_height) = img.dimensions();
    let img = img.resize(img_width / 25, img_height / 25, FilterType::Triangle)
                 .resize(img_width, img_height, FilterType::Nearest);
    /*
    let img = File::open(&args[1])?;
    let img = jpeg::JPEGDecoder::new(img);
    let img_buf: Vec<u8> = img.read_image().unwrap().into();
    let img = image::ImageBuffer::from_raw(img_width, img_height, img_buf.as_slice()).unwrap();
    */

    let buf = File::create(&args[2])?;
    let out_img = png::PNGEncoder::new(buf);
    //let (img_width, img_height) = img.dimensions();
    out_img.encode(&(img.raw_pixels()), img_width, img_height, ColorType::RGB(8))?;

    Ok(())
}
