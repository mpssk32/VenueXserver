CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    sender_id UUID REFERENCES users(id) ON DELETE CASCADE,

    receiver_id UUID REFERENCES users(id) ON DELETE CASCADE,

    content TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);