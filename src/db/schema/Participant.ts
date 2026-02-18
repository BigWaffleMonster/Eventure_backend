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
import { event } from './Event'

export const participant = pgTable('participant', {
  id: varchar('id')
    .$defaultFn(() => createId())
    .primaryKey(),

  // status: varchar({ enum: ['Yes', "No"] }).notNull(),
  dateCreated: timestamp({ withTimezone: true }).defaultNow(),
  dateUpdated: timestamp({ withTimezone: true }).defaultNow(),

  user_id: varchar('user_id')
    .notNull()
    .unique()
    .references(() => users.id),
  eventId: varchar('event_id')
    .unique()
    .notNull()
    .references(() => event.id)
})

export const participantTable = {
  participant
} as const

export type ParticipantTable = typeof participantTable
