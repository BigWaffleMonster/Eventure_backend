import { pgTable, varchar, timestamp, boolean } from 'drizzle-orm/pg-core'

import { createId } from '@paralleldrive/cuid2'

export const users = pgTable('users', {
  id: varchar('id')
    .$defaultFn(() => createId())
    .primaryKey(),

  login: varchar('login').unique(),
  email: varchar('email').notNull().unique(),
  password: varchar().notNull(),

  isEmailConfirmed: boolean().default(false),

  dateCreated: timestamp({ withTimezone: true }).defaultNow(),
  dateUpdated: timestamp({ withTimezone: true }).defaultNow()
})

export const usersTable = {
  users
} as const

export type UserTable = typeof usersTable
