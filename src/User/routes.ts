import Elysia, { t, type Handler } from 'elysia'
import { profile } from './controller/profile.controller'
import { signIn } from './controller/signIn.controller'
import { signUp } from './controller/signUp.controller'
import { initDB } from '../db/initDB'
import { signUpType } from './types'
import jwt from '@elysiajs/jwt'

export const userRouter = new Elysia({ prefix: '/user' }).get(
  '/profile',
  profile
)

export const authRouter = new Elysia({ prefix: '/auth' })
  .decorate('db', initDB())
  .use(
    jwt({
      name: 'accessToken',
      secret: 'superSecrete',
      exp: '30m'
    })
  )
  .use(
    jwt({
      name: 'refreshToken',
      secret: 'superSecreteForRefresh',
      exp: '7d'
    })
  )

  .post('/sign-in', ({ db, accessToken, refreshToken, cookie: { auth } }) => {
    return 'LOL'
  })

  .post(
    '/sign-up',
    async ({ body, db }) => {
      return await signUp({ body, db })
    },
    {
      body: t.Object({
        email: t.String(),
        password: t.String()
      })
    }
  )
