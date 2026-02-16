import {
  pgTable,
  varchar,
  timestamp,
  integer,
  date,
  boolean
} from 'drizzle-orm/pg-core'

import { createId } from '@paralleldrive/cuid2'

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

  owner: varchar('owner').notNull().unique(),
  category: varchar('category').notNull().unique()
})

export const eventTable = {
  event
} as const

export type EventTable = typeof eventTable
