use rand::seq::SliceRandom;
use reqwest::StatusCode;

fn request(symbol: &str) -> Result<(), reqwest::Error> {
    // Request via GET
    // Wonderful StackOverflow thread about a seemingly simple question: https://stackoverflow.com/questions/30154541/how-do-i-concatenate-strings
    let mut res =
        reqwest::blocking::get(&["http://localhost:3000/stocks?symbol=", symbol].concat())?;

    println!("Status: {}", res.status());

    println!("Headers:\n{:#?}", res.headers());
    // copy the response body directly to stdout
    res.copy_to(&mut std::io::stdout())?;

    Ok(())
}

fn main() -> () {
    // Define symbols
    let symbols = vec!["PEAR", "SQRL", "NULL", "ZVZZT"];
    // Do this 1,000 times
    let mut i = 0;
    while i < 10000 {
        // Chose a random element, returns Option<&Self::Item>
        let default = "PEAR";
        let symbol = symbols.choose(&mut rand::thread_rng()).unwrap_or(&default);
        println!("Testing API with {}", symbol);
        let res = request(symbol);
        // this will crash it once we crash the server
        println!("{:?}", res.unwrap());
        i = i + 1;
    }
}
