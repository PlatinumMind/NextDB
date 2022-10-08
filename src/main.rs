mod database;
use database::{KarmaDatabase, Kql};

fn main() {
    let db = KarmaDatabase::new("my-database", "root", "root");
}
