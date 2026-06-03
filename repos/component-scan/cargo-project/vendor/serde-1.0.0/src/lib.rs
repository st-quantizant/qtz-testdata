// serde: serialization/deserialization framework — no cryptography
use std::collections::HashMap;

pub trait Serialize {
    fn serialize(&self) -> String;
}

pub trait Deserialize: Sized {
    fn deserialize(s: &str) -> Option<Self>;
}

pub fn to_json<T: Serialize>(value: &T) -> String {
    value.serialize()
}
