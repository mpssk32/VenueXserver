CREATE TABLE applications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    artist_id UUID REFERENCES users(id) ON DELETE CASCADE,

    venue_id UUID REFERENCES venues(id) ON DELETE CASCADE,

    message TEXT NOT NULL,

    status TEXT DEFAULT 'pending',

    created_at TIMESTAMP DEFAULT NOW()
);