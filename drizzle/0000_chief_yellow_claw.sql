CREATE TABLE "category" (
	"id" varchar PRIMARY KEY NOT NULL,
	"title" varchar NOT NULL,
	CONSTRAINT "category_title_unique" UNIQUE("title")
);
--> statement-breakpoint
CREATE TABLE "event" (
	"id" varchar PRIMARY KEY NOT NULL,
	"title" varchar NOT NULL,
	"description" varchar NOT NULL,
	"participants" integer DEFAULT 0 NOT NULL,
	"maxParticipants" integer,
	"isOffline" boolean DEFAULT true,
	"location" varchar,
	"url" varchar,
	"startDate" date NOT NULL,
	"endDate" date,
	"dateCreated" timestamp with time zone DEFAULT now(),
	"dateUpdated" timestamp with time zone DEFAULT now(),
	"owner_id" varchar NOT NULL,
	"category_id" varchar NOT NULL
);
--> statement-breakpoint
CREATE TABLE "participant" (
	"id" varchar PRIMARY KEY NOT NULL,
	"dateCreated" timestamp with time zone DEFAULT now(),
	"dateUpdated" timestamp with time zone DEFAULT now(),
	"user_id" varchar NOT NULL,
	"event_id" varchar NOT NULL,
	CONSTRAINT "participant_user_id_unique" UNIQUE("user_id"),
	CONSTRAINT "participant_event_id_unique" UNIQUE("event_id")
);
--> statement-breakpoint
CREATE TABLE "user" (
	"id" varchar PRIMARY KEY NOT NULL,
	"login" varchar NOT NULL,
	"email" varchar NOT NULL,
	"password" varchar NOT NULL,
	"isEmailConfirmed" boolean DEFAULT false,
	"dateCreated" timestamp with time zone DEFAULT now(),
	"dateUpdated" timestamp with time zone DEFAULT now(),
	CONSTRAINT "user_login_unique" UNIQUE("login"),
	CONSTRAINT "user_email_unique" UNIQUE("email")
);
--> statement-breakpoint
ALTER TABLE "event" ADD CONSTRAINT "event_owner_id_user_id_fk" FOREIGN KEY ("owner_id") REFERENCES "public"."user"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "event" ADD CONSTRAINT "event_category_id_category_id_fk" FOREIGN KEY ("category_id") REFERENCES "public"."category"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "participant" ADD CONSTRAINT "participant_user_id_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."user"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "participant" ADD CONSTRAINT "participant_event_id_event_id_fk" FOREIGN KEY ("event_id") REFERENCES "public"."event"("id") ON DELETE no action ON UPDATE no action;