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
	Stroke()
	LineWidth(value int)
	ShadowBlur(value int)
	ShadowColor(value color.RGBA)
	ShadowOffsetX(value int)
	ShadowOffsetY(value int)
	AddColorStop(gradient interface{}, stop float64, color color.RGBA)
	CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{}
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{}
	FillStyle(value interface{})
	StrokeStyle(value interface{})
	GetImageData(x, y, width, height int) map[int]map[int]color.RGBA
	GetImageDataAlphaChannelOnly(x, y, width, height int) map[int]map[int]uint8
	GetImageDataHitByAlphaChannelValue(x, y, width, height int, minimumAcceptableValue uint8) map[int]map[int]bool
	ClearRect(x, y, width, height int)
	FillRect(x, y, width, height int)
}
