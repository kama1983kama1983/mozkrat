CREATE TABLE IF NOT EXISTS da3ormoz (

    id INTEGER PRIMARY KEY AUTOINCREMENT,

    title TEXT NOT NULL,

    Typeof TEXT CHECK(Typeof IN ('da3we', 'mozkrat')) NOT NULL,

    details TEXT,

    caseid INTEGER,

    notes TEXT,

    FOREIGN KEY (caseid) REFERENCES cases(id) ON DELETE SET NULL

);


CREATE TABLE IF NOT EXISTS qanon (

    id INTEGER PRIMARY KEY AUTOINCREMENT,

    title TEXT NOT NULL,

    date_of_publish DATE NOT NULL,

    mada text,

    FOREIGN KEY (mada) REFERENCES moad(mada) ON DELETE SET NULL

);


CREATE TABLE IF NOT EXISTS moad (

    id INTEGER PRIMARY KEY AUTOINCREMENT,

    mada TEXT,

    details TEXT,

    notes TEXT

);


CREATE TABLE IF NOT EXISTS al3aqod (

    id INTEGER PRIMARY KEY AUTOINCREMENT,

    title TEXT NOT NULL,

    details TEXT,

    date_contract DATE NOT NULL,

    first_side TEXT,

    secound_side TEXT,

    num_of_contract INTEGER,

    notes TEXT

);


CREATE TABLE IF NOT EXISTS cases (

    id INTEGER PRIMARY KEY AUTOINCREMENT,

    mod3 TEXT,

    mod3le TEXT,

    date_of_session DATE NOT NULL,

    tawkel TEXT

);