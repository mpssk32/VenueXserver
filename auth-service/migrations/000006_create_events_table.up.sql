CREATE TABLE events (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title TEXT NOT NULL,

    description TEXT,

    venue_id UUID REFERENCES venues(id) ON DELETE CASCADE,

    artist_id UUID REFERENCES users(id) ON DELETE CASCADE,

    event_date TIMESTAMP NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);