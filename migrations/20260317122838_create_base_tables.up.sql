CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    home_coordinates POINT NULLABLE,
    home_address TEXT NULLABLE,
    password VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    is_active BIT NOT NULL DEFAULT b'1',
    shift_start TIMESTAMP WITH TIME ZONE NULLABLE,
    shift_end TIMESTAMP WITH TIME ZONE NULLABLE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE user_teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL,
    team_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (team_id) REFERENCES teams (id)
);

CREATE TABLE vehicles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL,
    make VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    license_plate VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    vehicle_id UUID NOT NULL,
    status enum(
        'idle',
        'started',
        'completed',
        'cancelled'
    ) NOT NULL,
    pickup_time UUID NOT NULL,
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    trip_type enum('wth', 'htw',) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (vehicle_id) REFERENCES vehicles (id),
    FOREIGN KEY (pickup_time) REFERENCES time_slots (id)
);

CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL,
    token VARCHAR(255) NOT NULL,
    user_agent TEXT NULLABLE,
    ip_address INET,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    trip_id UUID NOT NULL,
    user_id UUID NOT NULL,
    status enum(
        'pending',
        'confirmed',
        'cancelled'
    ) NOT NULL,
    is_active BIT NOT NULL DEFAULT b'1',
    pickup_location TEXT,
    dropoff_location TEXT,
    pickup_coordinates POINT,
    dropoff_coordinates POINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (trip_id) REFERENCES trips (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE time_slots (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    time TIMESTAMP WITH TIME ZONE NOT NULL,
    is_active BIT NOT NULL DEFAULT b'1',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);