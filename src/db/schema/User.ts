import {
  pgTable,
  varchar,
  timestamp,
  integer,
  date,
  boolean
} from 'drizzle-orm/pg-core'

import { createId } from '@paralleldrive/cuid2'

export const user = pgTable('user', {
  id: varchar('id')
    .$defaultFn(() => createId())
    .primaryKey(),

  username: varchar('username').notNull().unique(),
  email: varchar('email').notNull().unique(),

  password: varchar().notNull(),

  dateCreated: timestamp({ withTimezone: true }).defaultNow(),
  dateUpdated: timestamp({ withTimezone: true }).defaultNow(),

  isEmailConfirmed: boolean().default(false)
})

export const userTable = {
  user
} as const

export type UserTable = typeof userTable

// type User struct {
// 	ID               uuid.UUID `gorm:"primaryKey;type:uuid"`
// 	UserName         string    `gorm:"index:unique"`
// 	Email            string    `gorm:"index:unique;not null"`
// 	Password         string
// 	DateCreated      time.Time
// 	DateUpdated      time.Time `gorm:"autoUpdateTime"`
// 	IsEmailConfirmed bool      `gorm:"default:false"`
// }
