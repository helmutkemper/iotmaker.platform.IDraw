package iotmaker_platform_IDraw

import "image/color"

// en: Interface with drawing methods. Originally this interface was based on web browsers and the canvas element its
// documentation was based on w3school and https://developer.mozilla.org/
//
// pt_br: Interface com os métodos de desenho. Originalmente esta interface teve como base os navegadores web e o
// elemento canvas sua documentação foi baseada no w3school e o site https://developer.mozilla.org/
type IDraw interface {

	// en: Begins a path, or resets the current path
	//     Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and arc(), to create paths.
	//     Tip: Use the stroke() method to actually draw the path on the canvas.
	//
	// pt_br: Inicia ou reinicializa uma nova rota no desenho
	//     Dica: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), e arc(), para criar uma nova rota no desenho
	//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
	BeginPath()

	// en: Moves the path to the specified point in the canvas, without creating a line
	//     x: The x-coordinate of where to move the path to
	//     y: The y-coordinate of where to move the path to
	//     Tip: Use the stroke() method to actually draw the path on the canvas.
	//
	// pt_br: Move o caminho do desenho para o ponto dentro do elemento canvas, sem inicializar uma linha
	//     X: Coordenada x para onde o ponto vai ser deslocado
	//     Y: Coordenada y para onde o ponto vai ser deslocado
	//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
	MoveTo(x, y int)

	// en: Creates an arc/curve between two tangents
	//     x0:     The x-axis coordinate of the first control point.
	//     y0:     The y-axis coordinate of the first control point.
	//     x1:     The x-axis coordinate of the second control point.
	//     y1:     The y-axis coordinate of the second control point.
	//     radius: The arc's radius. Must be non-negative.
	//
	// pt_br: Cria um arco/curva entre duas tangentes
	//     x0:     Eixo x da primeira coordenada de controle
	//     y0:     Eixo y da primeira coordenada de controle
	//     x1:     Eixo x da segunda coordenada de controle
	//     y1:     Eixo y da segunda coordenada de controle
	//
	//     Example:
	//     var c = document.getElementById("myCanvas");
	//     var ctx = c.getContext("2d");
	//     ctx.beginPath();
	//     ctx.moveTo(20, 20);              // Create a starting point
	//     ctx.lineTo(100, 20);             // Create a horizontal line
	//     ctx.arcTo(150, 20, 150, 70, 50); // Create an arc
	//     ctx.lineTo(150, 120);            // Continue with vertical line
	//     ctx.stroke();                    // Draw it
	ArcTo(x, y, radius, startAngle, endAngle int)

	// en: Adds a new point and creates a line from that point to the last specified point in the canvas. (this method does not draw the line).
	//     x: The x-coordinate of where to create the line to
	//     y: The y-coordinate of where to create the line to
	//     Tip: Use the stroke() method to actually draw the path on the canvas.
	//
	// pt_br: Adiciona um novo ponto e cria uma linha ligando o ponto ao último ponto especificado no elemento canvas. (este método não desenha uma linha).
	//     x: coordenada x para a criação da linha
	//     y: coordenada y para a criação da linha
	//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
	LineTo(x, y int)

	// en: Creates a path from the current point back to the starting point
	//     Tip: Use the stroke() method to actually draw the path on the canvas.
	//     Tip: Use the fill() method to fill the drawing (black is default). Use the fillStyle property to fill with
	//     another color/gradient.
	//
	// pt_br: cria um caminho entre o último ponto especificado e o primeiro ponto
	//     Dica: Use o método stroke() para desenhar a rota no elemento canvas
	//     Dica: Use o método fill() para preencher o desenho (petro é a cor padrão). Use a propriedade fillStyle para mudar
	//     a cor de preenchimento ou adicionar um gradiente
	ClosePath(x, y int)

	// en: The stroke() method actually draws the path you have defined with all those moveTo() and lineTo() methods. The default color is black.
	//     Tip: Use the strokeStyle property to draw with another color/gradient.
	//
	// pt_br: O método stroke() desenha o caminho definido com os métodos moveTo() e lineTo() usando a cor padrão, preta.
	//     Dica: Use a propriedade strokeStyle para desenhar com outra cor ou usar um gradiente
	Stroke()

	// en: Sets the current line width in pixels
	//     Default value: 1
	//     JavaScript syntax: context.lineWidth = number;
	//
	// pt_br: Define a espessura da linha em pixels
	//     Valor padrão: 1
	//     Sintaxe JavaScript: context.lineWidth = número
	//
	//     Example:
	//     var c = document.getElementById("myCanvas");
	//     var ctx = c.getContext("2d");
	//     ctx.lineWidth = 10;
	//     ctx.strokeRect(20, 20, 80, 100);
	SetLineWidth(value int)

	// en: Return the current line width in pixels
	//     Default value: 1
	//     JavaScript syntax: var l = context.lineWidth;
	//
	// pt_br: Retorna a espessura da linha em pixels
	//     Valor padrão: 1
	//     Sintaxe JavaScript: var l = context.lineWidth;
	//
	//     Example:
	//     var c = document.getElementById("myCanvas");
	//     var ctx = c.getContext("2d");
	//     ctx.lineWidth = 10;
	//     ctx.strokeRect(20, 20, 80, 100);
	//     var l = ctx.lineWidth;
	GetLineWidth() int

	// en: Sets the blur level for shadows
	//     Default value: 0
	//
	// pt_br: Define o valor de borrão da sombra
	//     Valor padrão: 0
	SetShadowBlur(value int)

	// en: Return the blur level for shadows
	//     Default value: 0
	//
	// pt_br: Retorna o valor de borrão da sombra
	//     Valor padrão: 0
	GetShadowBlur() int

	// en: Sets the color to use for shadows
	//     Note: Use the shadowColor property together with the shadowBlur property to create a shadow.
	//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY properties.
	//     Default value: #000000
	//
	// pt_br: Define a cor da sombra
	//     Nota: Use a propriedade shadowColor em conjunto com a propriedade shadowBlur para criar a sombra
	//     Dica: Ajuste o local da sombra usando as propriedades shadowOffsetX e shadowOffsetY
	//     Valor padrão: #000000
	SetShadowColor(value color.RGBA)

	// en: Sets the horizontal distance of the shadow from the shape
	//     shadowOffsetX = 0 indicates that the shadow is right behind the shape.
	//     shadowOffsetX = 20 indicates that the shadow starts 20 pixels to the right (from the shape's left position).
	//     shadowOffsetX = -20 indicates that the shadow starts 20 pixels to the left (from the shape's left position).
	//     Tip: To adjust the vertical distance of the shadow from the shape, use the shadowOffsetY property.
	//     Default value: 0
	//
	// pt_br: Define a distância horizontal entre a forma e a sua sombra
	//     shadowOffsetX = 0 indica que a forma e sua sombra estão alinhadas uma em cima da outra.
	//     shadowOffsetX = 20 indica que a forma e a sua sombra estão 20 pixels afastadas a direita (em relação a parte mais a esquerda da forma)
	//     shadowOffsetX = -20 indica que a forma e a sua sombra estão 20 pixels afastadas a esquerda (em relação a parte mais a esquerda da forma)
	//     Dica: Para ajustar a distância vertical, use a propriedade shadowOffsetY
	//     Valor padrão: 0
	ShadowOffsetX(value int)

	// en: Sets or returns the vertical distance of the shadow from the shape
	//     The shadowOffsetY property sets or returns the vertical distance of the shadow from the shape.
	//     shadowOffsetY = 0 indicates that the shadow is right behind the shape.
	//     shadowOffsetY = 20 indicates that the shadow starts 20 pixels below the shape's top position.
	//     shadowOffsetY = -20 indicates that the shadow starts 20 pixels above the shape's top position.
	//     Tip: To adjust the horizontal distance of the shadow from the shape, use the shadowOffsetX property.
	//     Default value: 0
	//
	// pt_br: Define a distância vertical entre a forma e a sua sombra
	//     shadowOffsetY = 0 indica que a forma e sua sombra estão alinhadas uma em cima da outra.
	//     shadowOffsetY = 20 indica que a forma e a sua sombra estão 20 pixels afastadas para baixo (em relação a parte mais elevada da forma)
	//     shadowOffsetY = -20 indica que a forma e a sua sombra estão 20 pixels afastadas para cima (em relação a parte mais elevada da forma)
	//     Dica: Para ajustar a distância horizontal, use a propriedade shadowOffsetX
	//     Valor padrão: 0
	ShadowOffsetY(value int)

	// en: Specifies the colors and stop positions in a gradient object
	//     gradient:     A gradient object created by CreateLinearGradient() or CreateRadialGradient() methods
	//     stopPosition: A value between 0.0 and 1.0 that represents the position between start (0%) and end (100%) in a gradient
	//     color:        A color RGBA value to display at the stop position
	//
	//     Note: You can call the addColorStopPosition() method multiple times to change a gradient. If you omit this method for
	//     gradient objects, the gradient will not be visible. You need to create at least one color stop to have a visible
	//     gradient.
	//
	// pt_br: Especifica a cor e a posição final para a cor dentro do gradiente
	//     gradient:     Objeto de gradiente criado pelos métodos CreateLinearGradient() ou CreateRadialGradient()
	//     stopPosition: Um valor entre 0.0 e 1.0 que representa a posição entre o início (0%) e o fim (100%) dentro do gradiente
	//     color:        Uma cor no formato RGBA para ser mostrada na posição determinada
	//
	//     Nota: Você pode chamar o método AddColorStopPosition() várias vezes para adicionar várias cores ao gradiente, porém, se
	//     você omitir o método, o gradiente não será visivel. Você tem a obrigação de chamar o método pelo menos uma vez com uma cor
	//     para que o gradiente seja visível.
	AddColorStopPosition(gradient interface{}, stop float64, color color.RGBA)

	// en: The fill() method fills the current drawing (path). The default color is black.
	//     Tip: Use the fillStyle property to fill with another color/gradient.
	//     Note: If the path is not closed, the fill() method will add a line from the last point to the start point of the
	//     path to close the path (like closePath()), and then fill the path.
	// pt_br: O método fill() preenche e pinta do desenho (caminho). A cor padrão é preto.
	//     Dica: Use a propriedade fillStyle para adicionar uma cor ou gradient.
	//     Nota: Se o caminho não estiver fechado, o método fill() irá adicioná uma linha do último ao primeiro ponto do
	//     caminho para fechar o caminho (semelhante ao método closePath()) e só então irá pintar
	Fill()

	// en: This method of the Canvas 2D API creates a gradient along the line connecting two given coordinates, starting at (x0, y0) point and ending at (x1, y1) point
	//     x0: The x-coordinate of the start point of the gradient
	//     y0: The y-coordinate of the start point of the gradient
	//     x1: The x-coordinate of the end point of the gradient
	//     y1: The y-coordinate of the end point of the gradient
	//
	//     The createLinearGradient() method creates a linear gradient object for be used with methods AddColorStopPosition(), FillStyle() and StrokeStyle().
	//     The gradient can be used to fill rectangles, circles, lines, text, etc.
	//     Tip: Use this object as the value to the strokeStyle() or fillStyle() methods
	//     Tip: Use the addColorStopPosition() method to specify different colors, and where to position the colors in the gradient object.
	//
	// pt_br: Este método do canvas 2D cria um gradiente ao longo de uma linha conectando dois pontos, iniciando no ponto (x0, y0) e terminando no ponto (x1, y1)
	//     x0: Coordenada x do ponto inicial do gradiente
	//     y0: Coordenada y do ponto inicial do gradiente
	//     x1: Coordenada x do ponto final do gradiente
	//     y1: Coordenada y do ponto final do gradiente
	//
	//     O método CreateLinearGradient() cria um objeto de gradiente linear para ser usado em conjunto com os métodos AddColorStopPosition(), FillStyle() e StrokeStyle().
	//     O gradiente pode ser usado para preencher retângulos, circulos, linhas, textos, etc.
	//     Dica: Use este objeto como valor passados aos métodos strokeStyle() ou fillStyle()
	//     Dica: Use o método addColorStopPosition() para especificar diferentes cores para o gradiente e a posição de cada cor
	CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{}

	// en: Creates a radial gradient (to use on canvas content). The parameters represent two circles, one with its center at (x0, y0) and a radius of r0, and the other with its center at (x1, y1) with a radius of r1.
	//     x0: The x-coordinate of the starting circle of the gradient
	//     y0: The y-coordinate of the starting circle of the gradient
	//     r0: The radius of the starting circle. Must be non-negative and finite. (note: radius is a width, not a degrees angle)
	//     x1: The x-coordinate of the ending circle of the gradient
	//     y1: The y-coordinate of the ending circle of the gradient
	//     r1: The radius of the ending circle. Must be non-negative and finite. (note: radius is a width, not a degrees angle)
	//
	// pt_br: Este método cria um gradiente radial (para ser usado com o canvas 2D). Os parâmetros representam dois círculos, um com o centro no ponto (x0, y0) e raio r0, e outro com centro no ponto (x1, y1) com raio r1
	//     x0: Coordenada x do circulo inicial do gradiente
	//     y0: Coordenada y do circulo inicial do gradiente
	//     r0: Raio do círculo inicial. Deve ser um valor positivo e finito. (nota: o raio é um comprimento e não um ângulo)
	//     x1: Coordenada x do circulo final do gradiente
	//     y1: Coordenada y do circulo final do gradiente
	//     r1: Raio do círculo final. Deve ser um valor positivo e finito. (nota: o raio é um comprimento e não um ângulo)
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{}
	FillStyle(value interface{})
	StrokeStyle(value interface{})
	GetImageData(x, y, width, height int) map[int]map[int]color.RGBA
	GetImageDataAlphaChannelOnly(x, y, width, height int) map[int]map[int]uint8
	GetImageDataCollisionByAlphaChannelValue(x, y, width, height int, minimumAcceptableValue uint8) map[int]map[int]bool
	ClearRect(x, y, width, height int)
	FillRect(x, y, width, height int)
}
