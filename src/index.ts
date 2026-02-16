import 'dotenv/config'
import cors from '@elysiajs/cors'
import jwt from '@elysiajs/jwt'
import { drizzle } from 'drizzle-orm/singlestore/driver'
import { Elysia } from 'elysia'

const app = new Elysia()

app.use(
  cors({
    credentials: true,
    origin: '*'
  })
)

app.use(
  jwt({
    name: 'accessToken',
    secret: 'superSecrete',
    exp: '30m'
  })
)

app.use(
  jwt({
    name: 'refreshToken',
    secret: 'superSecreteForRefresh',
    exp: '7d'
  })
)

app.get('/', () => 'Hello Elysia').listen(3000)

try {
  const db = drizzle(process.env.DB_URL!)
} catch (e) {
  console.log(e)
}

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
)
