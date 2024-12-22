-- Create the UserRole table if it doesn't already exist
CREATE TABLE IF NOT EXISTS public.UserRole (
  id uuid PRIMARY KEY REFERENCES auth.users(id),
  role text NOT NULL DEFAULT 'member'
);

-- Create the trigger function
CREATE OR REPLACE FUNCTION create_user_role()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO public.UserRole (id, role)
  VALUES (NEW.id, 'member');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER after_user_created
AFTER INSERT ON auth.users
FOR EACH ROW
EXECUTE FUNCTION create_user_role();
