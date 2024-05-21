package constants

const DatabasePath = "./reach.db"

const SqlPrep = `
CREATE TABLE IF NOT EXISTS Action (
    id INTEGER PRIMARY KEY,
    status INTEGER NOT NULL,
    desc TEXT NOT NULL,
    difficulty INTEGER NOT NULL,
    notes TEXT NOT NULL,
    start_strategy TEXT NOT NULL,
    reflection TEXT NOT NULL,
    outcome INTEGER NOT NULL
);
`
