package cmd

import "context"

var srv server

func load(ctx context.Context) error {

	if err := srv.loadConfig(ctx); err != nil {
		return err
	}

	if err := srv.loadLogger(); err != nil {
		return err
	}

	if err := srv.loadDatabaseClients(ctx); err != nil {
		return err
	}

	if err := srv.loadThirdPartyClients(ctx); err != nil {
		return err
	}

	if err := srv.loadRepositories(); err != nil {
		return err
	}

	if err := srv.loadServices(); err != nil {
		return err
	}

	if err := srv.loadDeliveries(); err != nil {
		return err
	}

	if err := srv.loadPublishers(ctx); err != nil {
		return err
	}

	if err := srv.loadSubscribers(ctx); err != nil {
		return err
	}

	if err := srv.loadClients(ctx); err != nil {
		return err
	}

	if err := srv.loadServers(ctx); err != nil {
		return err
	}

	return nil
}
func start(ctx context.Context) error {
	errChan := make(chan error)

	for _, database := range srv.factories {
		if err := database.Connect(ctx); err != nil {
			return err
		}
	}

	for _, p := range srv.processors {
		go func(p processor) {
			if err := p.Start(ctx); err != nil {
				errChan <- err
			}
		}(p)
	}
	go func() {
		err := <-errChan
		srv.logger.Sugar().Errorf("start error: %w\n", err)
	}()
	return nil
}

func stop(ctx context.Context) error {
	for _, processor := range srv.processors {
		if err := processor.Stop(ctx); err != nil {
			return err
		}
	}

	for _, database := range srv.factories {
		if err := database.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}
