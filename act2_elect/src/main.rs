use std::{io, usize};

// 2.) Iterates over the array and converts each number into the chosen number system of the user.
fn convert_arr(arr: &mut Vec<String>, i: usize) {
    for element in arr.iter_mut() {
        *element = match element.parse::<usize>() {
            Ok(num) => {
                match i {
                    1 => format!("{:b}", num),
                    2 => format!("{:o}", num),
                    3 => format!("{:X}", num),
                    _ => {
                        println!("error converting number");
                        continue;
                    },
                }
            },
            Err(_) => {
                println!("Invalid number: {}. Can't be converted into decimal.", element);
                continue;
            }
        }
    }
}

// 3.) Iterates over the array and converts the values into their corresponding ASCII values.
fn to_ascii(arr: &mut Vec<String>) {
    for element in arr.iter_mut() {
        match element.parse::<u32>() {
            Ok(num) => {
                if let Some(ch) = char::from_u32(num) {
                    *element = ch.to_string();
                } else {
                    println!("invalid ascii code")
                }
            },
            Err(_) => println!("invalid number"),
        }
    }
}

fn main() {
    // 1.) Asks for user input for the array size. Will ask for input again if user input is invalid.
    let size: usize = loop {
        let mut input = String::new();
        println!("How many numbers will you enter: ");

        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line.");

        match input.trim().parse() {
            Ok(num) => break num,
            Err(_) => {
                println!("Please enter an integer.");
                continue
            },
        };
    };

    // 1.) Asks the user for the number system to convert to. Asks again if useer input is invalid.
    let num_system: usize = loop {
        let mut input = String::new();
        println!("");
        println!("Choose a number system to convert to:");
        println!("1. Binary");
        println!("2. Octal");
        println!("3. Hexadecimal");

        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line.");

        match input.trim().parse() {
            Ok(num) if (1..=3).contains(&num) => break num,
            Ok(_) => println!("Invalid choice. Choose only 1, 2, or 3."),
            Err(_) => {
                println!("Please enter an integer.");
                continue
            },
        };
    };

    // 2.) Enter numbers for each index. 
    let mut arr: Vec<String> = vec![String::new(); size];
    
    for (i, element) in arr.iter_mut().enumerate() {
        loop {
            println!("");
            println!("Enter a number ({}/{}): ", i+1, size);

            let mut num = String::new();
            io::stdin()
                .read_line(&mut num)
                .expect("Failed to read line.");

            match num.trim().parse::<usize>() {
                Ok(num) => {
                    *element = num.to_string();
                    break;
                },
                Err(_) => println!("Please enter an integer.")
            }
        }
    }

    println!("");
    println!("Array before conversion: {:?}", arr);
    
    // 2.) Convert each value in the array to the chosen number system of the user.
    let mut conv_arr = arr.clone();
    convert_arr(&mut conv_arr, num_system);
    println!("Array after conversion: {:?}", conv_arr);
    
    // 3.) Convert each value in the array to its corresponding ASCII value.
    to_ascii(&mut arr);
    println!("ASCII values: {:?}", arr);

}