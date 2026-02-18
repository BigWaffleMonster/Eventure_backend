// src/db/index.ts
import 'dotenv/config'
import { drizzle } from 'drizzle-orm/postgres-js'
import { sql } from 'drizzle-orm'

let _db: ReturnType<typeof drizzle> | null = null

export function initDB() {
  if (!_db) {
    if (!process.env.DB_URL) {
      throw new Error('❌ DB_URL is required')
    }
    try {
      _db = drizzle(process.env.DB_URL)
      console.log('✅ Database initialized')
    } catch (e) {
      throw e
    }
  }
  return _db
}

export type DBType = ReturnType<typeof initDB>

export async function checkDBConnection(): Promise<void> {
  try {
    const db = initDB()
    await db.execute('select 1')
    console.log('✅ Database connection verified')
  } catch (error) {
    console.error('❌ Database connection failed:', error)
    throw error
  }
}
