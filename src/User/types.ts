import { DBType } from '../db/initDB'

export type signUpType = {
  body: { password: string; email: string }
  db: DBType
}
