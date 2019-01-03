// A decorating helper to wrap Confita's backends with an error-suppressing logic. confita.Backend and confita.Unmarshaller interfaces are supported.
//
// By default Confita loader fails as soon as one of backends returns an error. In some cases these errors (e.g.: optional non-existing local configuration file or non-accessible dev instance of Consul) should not be considered as critical ones. Since Confita doesn't provide a strategy to handle this case of out the box, such backends could be decorated to suppress occurred errors.
package maybe
