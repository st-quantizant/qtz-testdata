// clap: command-line argument parser — no cryptography
use std::collections::HashMap;
use std::env;

pub struct Command {
    name: String,
    args: Vec<Arg>,
}

pub struct Arg {
    name: String,
    required: bool,
    default: Option<String>,
}

impl Command {
    pub fn new(name: &str) -> Self {
        Self { name: name.to_string(), args: vec![] }
    }

    pub fn arg(mut self, arg: Arg) -> Self {
        self.args.push(arg);
        self
    }

    pub fn get_matches(self) -> HashMap<String, String> {
        let raw: Vec<String> = env::args().skip(1).collect();
        let mut matches = HashMap::new();
        for (i, a) in self.args.iter().enumerate() {
            if let Some(v) = raw.get(i) {
                matches.insert(a.name.clone(), v.clone());
            } else if let Some(d) = &a.default {
                matches.insert(a.name.clone(), d.clone());
            }
        }
        matches
    }
}

impl Arg {
    pub fn new(name: &str) -> Self {
        Self { name: name.to_string(), required: false, default: None }
    }

    pub fn required(mut self, r: bool) -> Self { self.required = r; self }
    pub fn default_value(mut self, v: &str) -> Self { self.default = Some(v.to_string()); self }
}
