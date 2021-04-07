import { verify } from 'jsonwebtoken'

interface Token {
  access_token: string
  expires_in: number
  scope: string
  token_type: string
}

const { JWT_SECRET } = process.env

export function verifyToken(token: string): Token | undefined {
  if (!JWT_SECRET) throw new Error('[Jsonwebtoken] Missing JWT_SECRET')

  try {
    return (verify(token, JWT_SECRET) as any) || undefined
  } catch {
    return undefined
  }
}
