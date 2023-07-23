CREATE TABLE "block_headers" (
  "slot" bigInt PRIMARY KEY NOT NULL,
  "root" varchar NOT NULL,
  "canonical" bool NOT NULL,
  "proposer_index" bigInt NOT NULL,
  "parent_root" varchar NOT NULL,
  "state_root" varchar NOT NULL,
  "body_root" varchar NOT NULL,
  "signature" varchar NOT NULL,
  "finalized" bool NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "attestation_duties" (
  "id" BIGSERIAL PRIMARY KEY,
  "pubkey" varchar NOT NULL,
  "validator_index" bigInt NOT NULL,
  "committees_at_slot" bigInt NOT NULL,
  "committee_index" bigInt NOT NULL,
  "committee_length" bigInt NOT NULL,
  "validator_committee_index" bigInt NOT NULL,
  "slot" bigInt NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
 
CREATE TABLE "attestation_details" (
  "id" BIGSERIAL PRIMARY KEY,
  "finalized" bool NOT NULL,
  "aggregation_bits" varchar NOT NULL,
  "slot" bigInt NOT NULL,
  "index" bigInt NOT NULL,
  "beacon_block_root" varchar NOT NULL,
  "source_epoch" bigInt NOT NULL,
  "source_root" varchar NOT NULL,
  "target_epoch" bigInt NOT NULL,
  "target_root" varchar NOT NULL,
  "signature" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "validators" (
    "index" bigInt PRIMARY KEY,
    "balance" bigInt NOT NULL,
    "status" varchar NOT NULL,
    "pubkey" varchar NOT NULL,
    "withdrawal_credentials" varchar NOT NULL,
    "effective_balance" bigInt NOT NULL,
    "slashed" bool NOT NULL,
    "activation_eligibility_epoch" varchar NOT NULL,
    "activation_epoch" varchar NOT NULL,
    "exit_epoch" varchar NOT NULL,
    "withdrawable_epoch" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "block_headers" ("root");

CREATE INDEX ON "block_headers" ("proposer_index");

CREATE INDEX ON "attestation_duties" ("pubkey");

CREATE INDEX ON "attestation_duties" ("validator_index");

CREATE INDEX ON "attestation_duties" ("committee_index");

CREATE INDEX ON "attestation_duties" ("slot");

CREATE INDEX ON "attestation_details" ("slot");

CREATE INDEX ON "validators" ("pubkey");


