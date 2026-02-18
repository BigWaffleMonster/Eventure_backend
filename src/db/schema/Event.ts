import {
  pgTable,
  varchar,
  timestamp,
  integer,
  date,
  boolean
} from 'drizzle-orm/pg-core'

import { createId } from '@paralleldrive/cuid2'
import { users } from './Users'
import { category } from './Category'

export const event = pgTable('event', {
  id: varchar('id')
    .$defaultFn(() => createId())
    .primaryKey(),

  title: varchar('title').notNull(),
  description: varchar('description').notNull(),

  participants: integer().default(0).notNull(),
  maxParticipants: integer(),

  isOffline: boolean().default(true),
  location: varchar('location'),
  url: varchar('url'),

  startDate: date().notNull(),
  endDate: date(),

  dateCreated: timestamp({ withTimezone: true }).defaultNow(),
  dateUpdated: timestamp({ withTimezone: true }).defaultNow(),

  ownerId: varchar('owner_id')
    .notNull()
    .references(() => users.id),
  categoryId: varchar('category_id')
    .notNull()
    .references(() => category.id)
})

export const eventTable = {
  event
} as const

export type EventTable = typeof eventTable
