package config

type Supabase struct {
	JWTSecret string
}

func LoadConfig() Supabase {
	return Supabase{
		JWTSecret: "TODO",
	}
}
