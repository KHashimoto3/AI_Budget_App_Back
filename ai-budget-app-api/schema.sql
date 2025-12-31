-- public schema

-- users
CREATE TABLE IF NOT EXISTS public.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firebase_uid TEXT NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    disp_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    account_type INTEGER NOT NULL DEFAULT 0 CHECK (account_type IN (0, 1)),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    login_at TIMESTAMP,
    deleted_flag BOOLEAN NOT NULL DEFAULT FALSE
);

-- genres
CREATE TABLE IF NOT EXISTS public.genres (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL
);

-- images
CREATE TABLE IF NOT EXISTS public.images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    filename VARCHAR NOT NULL,
    file_path VARCHAR NOT NULL,
    parsed_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- expenses
CREATE TABLE IF NOT EXISTS public.expenses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    expense_date DATE NOT NULL,
    amount INTEGER NOT NULL CHECK (amount >= 0),
    genres_id UUID REFERENCES public.genres(id) ON DELETE SET NULL,
    shop_name VARCHAR NOT NULL,
    memo TEXT,
    input_type VARCHAR NOT NULL CHECK (input_type IN ('manual', 'ocr')),
    image_id UUID REFERENCES public.images(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_users_firebase_uid ON public.users(firebase_uid);
CREATE INDEX IF NOT EXISTS idx_users_email ON public.users(email);
CREATE INDEX IF NOT EXISTS idx_expenses_user_id ON public.expenses(user_id);
CREATE INDEX IF NOT EXISTS idx_expenses_expense_date ON public.expenses(expense_date);
CREATE INDEX IF NOT EXISTS idx_expenses_genres_id ON public.expenses(genres_id);
CREATE INDEX IF NOT EXISTS idx_images_user_id ON public.images(user_id);
CREATE INDEX IF NOT EXISTS idx_genres_user_id ON public.genres(user_id);
