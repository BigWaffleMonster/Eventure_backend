import {
  pgTable,
  varchar,
  timestamp,
  integer,
  date,
  boolean
} from 'drizzle-orm/pg-core'

import { createId } from '@paralleldrive/cuid2'

export const category = pgTable('category', {
  id: varchar('id')
    .$defaultFn(() => createId())
    .primaryKey(),
  title: varchar('title').notNull().unique()
})

export const categoryTable = {
  category
} as const

export type CategoryTable = typeof categoryTable
