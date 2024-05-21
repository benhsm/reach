package constants

const DatabasePath = "./reach.db"

const SqlPrep = `
CREATE TABLE IF NOT EXISTS Action (
    id INTEGER NOT NULL PRIMARY KEY,
    status INTEGER NOT NULL,
    desc TEXT NOT NULL,
    difficulty INTEGER NOT NULL,
    notes TEXT NOT NULL,
    reflection TEXT NOT NULL,
    outcome_value INTEGER NOT NULL
);
`
