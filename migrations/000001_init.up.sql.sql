CREATE TABLE users (
    id BIGSERIAL primary key,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL unique ,
    password_hash varchar(255) NOT NULL ,
    role VARCHAR(50) NOT NULL default 'client',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()

    CONSTRAINT check_role CHECK ( role  IN ('admin' , 'client') )
);

CREATE TABLE membership_plans (
    id SERIAL primary key,
    name VARCHAR(255) NOT NULL ,
    price DECIMAL(10 , 2),
    duration_days INT NOT NULL ,
    visits_limit INT DEFAULT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT check_duration_positive CHECK ( duration_days > 0 ),
    CONSTRAINT check_price_positive CHECK (price >= 0)

);

CREATE TABLE user_memberships (
    id BIGSERIAL primary key,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE ,
    plan_id INT  NOT NULL REFERENCES membership_plans(id),
    start_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    end_date TIMESTAMPTZ NOT NULL ,
    visits_left INT DEFAULT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),


    CONSTRAINT check_dates CHECK ( end_date > start_date ),
    CONSTRAINT check_visits_not_negative CHECK ( visits_left >= 0 )

);

CREATE TABLE classes (
    id SERIAL primary key ,
    title VARCHAR(255) NOT NULL ,
    start_time TIMESTAMPTZ NOT NULL ,
    trainer_name VARCHAR(255) NOT NULL ,
    capacity INT NOT NULL,
    occupied INT NOT NULL DEFAULT 0,


    CONSTRAINT check_capacity_positive CHECK ( capacity >0 ),
    CONSTRAINT check_occupied_positive CHECK ( occupied >=0 )

);

CREATE TABLE bookings (
    id BIGSERIAL primary key ,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE ,
    class_id BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE ,
    status VARCHAR(100) NOT NULL DEFAULT 'reserved',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    unique (user_id, class_id)
);

CREATE INDEX idx_user_memberships_user_id ON user_memberships(user_id);
CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_classes_start_time ON classes(start_time);