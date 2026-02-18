import { Handler } from 'elysia'
import { signUpType } from '../types'
import { users } from '../../db/schema/Users'
import { initDB } from '../../db/initDB'
import { eq, sql } from 'drizzle-orm'

export const signUp = async ({ body, db }: signUpType) => {
  const { password, email } = body

  try {
    const existingUser = await db
      .select()
      .from(users)
      .where(eq(users.email, email))

    if (existingUser.length > 1) {
      console.log(existingUser, 'OIOILK')
      throw new Error('email exists')
    }

    const hasedPassword = await Bun.password.hash(password, {
      algorithm: 'bcrypt',
      cost: 8
    })

    await db.insert(users).values({ email, password: hasedPassword })

    return 'user created'
  } catch (e) {
    console.log(e)
  }
}
