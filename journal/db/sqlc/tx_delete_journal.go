package db

import "context"

func (store *SQLStore) DeleteJournalTx(ctx context.Context, journal_id int64) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		err = q.DeleteAudioOfJournal(ctx, journal_id)
		if err != nil {
			return err
		}

		err = q.DeleteImageOfJournal(ctx, journal_id)
		if err != nil {
			return err
		}

		err = q.DeleteJournal(ctx, journal_id)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}