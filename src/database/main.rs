use super::{types::KarmaDatabase, Kql};

impl KarmaDatabase {
    pub fn new(name: &str, username: &str, password: &str) -> Self {
        Self {
            name: name.to_string(),
            username: username.to_string(),
            password: password.to_string(),
            port: 9990,
        }
    }

    pub fn get(&self, key: &str) {
        todo!()
    }

    pub fn set(&self, key: &str, value: &str) {
        todo!()
    }

    pub fn execute(&self, kql_query: Kql) {
        todo!()
    }
}
