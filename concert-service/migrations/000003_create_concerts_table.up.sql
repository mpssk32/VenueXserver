CREATE TABLE concerts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    genre TEXT NOT NULL,
    concert_date TIMESTAMP NOT NULL,
    ticket_price INT NOT NULL,
    venue_id UUID REFERENCES venues(id) ON DELETE CASCADE,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);  