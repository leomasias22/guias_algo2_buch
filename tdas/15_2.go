package tdas

type MamushkaInterface interface {
	ObtenerColor() Color

	Guardar(aGuardar MamushkaInterface) bool

	ObtenerGuardada() MamushkaInterface
}
