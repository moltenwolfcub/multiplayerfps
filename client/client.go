package client

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/moltenwolfcub/gogl-utils"
	"github.com/moltenwolfcub/multiplayerfps/common"
	"github.com/veandco/go-sdl2/sdl"
)

type Client struct {
	listenAddr string
	connection common.Connection

	window        *sdl.Window
	keyboardState []uint8
	camera        *gogl.Camera

	shader  *gogl.Shader
	texture gogl.TextureID
	cube    gogl.Object

	lightingColor mgl32.Vec3
}

func NewClient(listenAddr string) *Client {
	return &Client{
		listenAddr: listenAddr,
	}
}

/*
Connects to the server and starts running the loops
which handle the rest of the logic
*/
func (c *Client) Start() error {
	conn, err := net.Dial("tcp", c.listenAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	c.connection = common.NewConnection(conn)

	cleanup := c.initialise()
	defer cleanup()

	go c.readLoop()
	return c.mainLoop()
}

/*
A loop to manage clientbound traffic and send recieved packets
to the handlepacket method for processing.
*/
func (c *Client) readLoop() error {
	for {
		rawPacket := c.connection.MustRecieve()
		err := c.handlePacket(rawPacket)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

// ONLY EDIT BELOW THIS LINE! The above code handles the client setup and manages the network connection

var (
	windowWidth  int32 = 1280
	windowHeight int32 = 720

	cameraFov  float32 = 45
	cameraNear float32 = 0.1
	cameraFar  float32 = 100.0
)

func (c *Client) initialise() func() {
	window, cleanup := gogl.SetupFPSWindow("Multiplayer FPS", windowWidth, windowHeight)
	c.window = window
	fmt.Println("OpenGL Version", gogl.GetVersion())

	window.WarpMouseInWindow(windowWidth/2, windowHeight/2)

	c.shader = gogl.NewShader("assets/shaders/test.vert", "assets/shaders/quadTexture.frag")
	c.texture = gogl.LoadTexture("assets/textures/metal/metalbox_full.png")

	c.cube = gogl.Cube(1)

	gl.BindVertexArray(0)

	c.connection.MustSend(common.ServerBoundLightingRequest{})

	c.keyboardState = sdl.GetKeyboardState()
	c.camera = gogl.NewCamera(mgl32.Vec3{}, mgl32.Vec3{0, 1, 0}, 0, 0, 0.0025, 0.1)

	return cleanup
}

/*
Main loop that'll handle the clientside logic and state.
*/
func (c *Client) mainLoop() error {

	elapsedTime := float32(0)
	for {
		frameStart := time.Now()

		//handleEvents
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return nil
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_RESIZED {
					windowWidth, windowHeight = e.Data1, e.Data2
					gl.Viewport(0, 0, windowWidth, windowHeight)
				}
			}
		}
		if c.keyboardState[sdl.SCANCODE_ESCAPE] != 0 {
			return nil
		}

		//updateCamera
		dirs := gogl.NewMoveDirs(
			c.keyboardState[sdl.SCANCODE_W] != 0,
			c.keyboardState[sdl.SCANCODE_S] != 0,
			c.keyboardState[sdl.SCANCODE_D] != 0,
			c.keyboardState[sdl.SCANCODE_A] != 0,
			c.keyboardState[sdl.SCANCODE_SPACE] != 0,
			c.keyboardState[sdl.SCANCODE_LSHIFT] != 0,
		)
		mouseX, mouseY, _ := sdl.GetMouseState()
		mouseDx, mouseDy := float32(mouseX-windowWidth/2), -float32(mouseY-windowHeight/2)
		c.camera.UpdateCamera(dirs, elapsedTime, mouseDx, mouseDy)
		// log.Println(mouseDx, mouseDy)

		//draw
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		c.shader.Use()
		projMat := mgl32.Perspective(mgl32.DegToRad(cameraFov), float32(windowWidth)/float32(windowHeight), cameraNear, cameraFar)
		viewMat := c.camera.GetViewMatrix()
		c.shader.SetMatrix4("proj", projMat)
		c.shader.SetMatrix4("view", viewMat)

		c.shader.SetVec3("viewPos", c.camera.Pos)
		c.shader.SetVec3("lightPos", mgl32.Vec3{3.3, 1, 0})
		c.shader.SetVec3("lightColor", c.lightingColor)
		c.shader.SetVec3("ambientLight", c.lightingColor.Mul(0.3))

		gogl.BindTexture(c.texture)

		c.cube.DrawMultiple(c.shader, 5, func(i int) mgl32.Mat4 {
			return mgl32.Ident4().Mul4(mgl32.Translate3D(1, 0, float32(2*i)))
		})

		//post draw
		c.window.GLSwap()
		c.shader.CheckShadersForChanges()
		elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)

		sdl.EventState(sdl.MOUSEMOTION, sdl.IGNORE)
		c.window.WarpMouseInWindow(windowWidth/2, windowHeight/2)
		sdl.EventState(sdl.MOUSEMOTION, sdl.ENABLE)
	}
}

/*
Will figure out what kind of packet has been recieved
and correctly handle how it should behave.
*/
func (c *Client) handlePacket(rawPacket common.Packet) error {
	switch packet := rawPacket.(type) {
	case common.ClientBoundLightingUpdate:
		c.lightingColor = packet.Color
	default:
		return fmt.Errorf("unkown packet: %s", packet)
	}
	return nil
}
