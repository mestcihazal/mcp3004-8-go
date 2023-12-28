// package main is a module that implements a sensor model of mcp300x
package main

import (
	"context"

	"github.com/mestcihazal/mcp3004-8-go/mcp300x"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("mcp3004-8-go"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	sensorModule, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	sensorModule.AddModelFromRegistry(ctx, sensor.API, mcp300x.Model)

	err = sensorModule.Start(ctx)
	defer sensorModule.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
