import 'dotenv/config'
import cors from '@elysiajs/cors'
import { Elysia } from 'elysia'
import { authRouter } from './User/routes'
import { checkDBConnection, initDB } from './db/initDB'

try {
  await checkDBConnection()
} catch (e) {
  process.exit(1)
}

const app = new Elysia()

app.use(
  cors({
    credentials: true,
    origin: '*'
  })
)

app.use(authRouter)

app.listen(3000)

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
)
