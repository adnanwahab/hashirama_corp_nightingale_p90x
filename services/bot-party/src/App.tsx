//import reactLogo from './assets/react.svg'
//import viteLogo from '/vite.svg'
//import { Joystick } from 'react-joystick-component';
import {useEffect, useState} from 'react'
//import ass from 'assets/most_recent.jpg'
import Parking from './Parking'
import React from 'react'
//import { loadPlayer } from 'rtsp-relay/browser';
/*<Joystick
   stickImage={'https://w7.pngwing.com/pngs/880/496/png-transparent-arrows-controller-direction-joystick-videogames-videogames-icon.png'}
   size={100} sticky={false} baseColor="steelblue" stickColor="aliceblue" move={handleMove} stop={handleStop}></Joystick> */


console.log('initializing app')
// const speech_1 = `
// Hello, and thank you for activating the Swarmbotics Unmanned ground Vehicle. You may call me Sheila.
// Would you like to start with a tutorial?
// `
// const speech_2 = `
//   okay. that sounds great
// `


// const content = [
//   speech_2,
//   speech_1
// ]

function speak() {
  // Create a SpeechSynthesisUtterance
  // const utterance = new SpeechSynthesisUtterance(content.pop() || 'thank you for driving. goodbye');

  // // Select a voice
  // const voices = speechSynthesis.getVoices();
  // utterance.voice = voices[0]; // Choose a specific voice

  // // Speak the text
  // speechSynthesis.speak(utterance);

  
}

speak();
console.log('hi')


// turn left 90 deg
// turn right 90 deg
// walk forward 90 feet
// walk backward 90 feet

//way point mode
//session replay
let base = `http://10.0.0.2:8000`
//base = `http://871f-72-195-220-81.ngrok-free.app`
base = `http://orin.local:8000`
base = `http://192.168.1.202:8000`

async function do_stuff (x:number, y:number, backend_url: string) {
  if (! x) x = Math.random()
  if (! y) y = Math.random()
  return('shit')
  //let url = backend_url || `https://3ab8-72-195-220-81.ngrok-free.app/drive`
  let url = `${base}/drive`

  //url = `https://8a70-72-195-220-81.ngrok-free.app/drive`

  //url = `https://192.168.1.221:8000/brake`

  try {
  let req = await fetch(url, {
    headers: {
      'Content-Type': 'application/json'
    },
    method: 'POST',
    //mode: 'no-cors',
    body: JSON.stringify({x, y})
  })
    let response = await req.json()
    console.log(response)

    console.log(response)
  } catch (e) {
    //console.log(e.message)
    //document.body.innerHTML = e.message + document.body.innerHTML
    /* document.body.innerHTML =

     *   new Array(1e6).fill(0).map(_ =>
     *     {
     *       return ('wow' +
     *        JSON.stringify(e))
     *     }

     *   ) */

  }
  //alert(1123, 'your mom')

  //debugger
}

async function send_brake_message () {
  let url = `${base}/brake`

  let req = await fetch(url, {
    headers: {
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({
      //x, y
    })
  })
  let response = await req.json()
  console.log(response)
}

// let one = `relative h-full w-full rounded-xl bg-white shadow-[0px_0px_0px_1px_rgba(9,9,11,0.07),0px_2px_2px_0px_rgba(9,9,11,0.05)] dark:bg-zinc-900 dark:shadow-[0px_0px_0px_1px_rgba(255,255,255,0.1)] dark:before:pointer-events-none dark:before:absolute dark:before:-inset-px dark:before:rounded-xl dark:before:shadow-[0px_2px_8px_0px_rgba(0,_0,_0,_0.20),_0px_1px_0px_0px_rgba(255,_255,_255,_0.06)_inset] forced-colors:outline`
// let two = `grid h-full w-full justify-items-center overflow-hidden place-items-center p-6 py-8 sm:p-8 lg:p-12`
// let three = `flex gap-x-6`

function getX(prevEvent: any) {
  return prevEvent.x?.toPrecision(2)
}
function getY(prevEvent: any) {
  return prevEvent.y?.toPrecision(2)
}



/* <a href="#_" class="relative inline-block px-4 py-2 font-medium group">
 * <span class="absolute inset-0 w-full h-full transition duration-200 ease-out transform translate-x-1 translate-y-1 bg-black group-hover:-translate-x-0 group-hover:-translate-y-0"></span>
 * <span class="absolute inset-0 w-full h-full bg-white border-2 border-black group-hover:bg-black"></span>
 * <span class="relative text-black group-hover:text-white">Button Text</span>
 * </a> */

/* <a href="#_" class="">
 * Button Text
 * </a> */

let four = `
hover:bg-indigo-500
focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600




w-16
ml-16
relative isolate inline-flex items-center justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(theme(spacing[3.5])-1px)] py-[calc(theme(spacing[2.5])-1px)] sm:px-[calc(theme(spacing.3)-1px)] sm:py-[calc(theme(spacing[1.5])-1px)] sm:text-sm/6 focus:outline-none data-[focus]:outline data-[focus]:outline-2 data-[focus]:outline-offset-2 data-[focus]:outline-blue-500 data-[disabled]:opacity-50 [&>[data-slot=icon]]:-mx-0.5 [&>[data-slot=icon]]:my-0.5 [&>[data-slot=icon]]:size-5 [&>[data-slot=icon]]:shrink-0 [&>[data-slot=icon]]:text-[--btn-icon] [&>[data-slot=icon]]:sm:my-1 [&>[data-slot=icon]]:sm:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-[hover]:[--btn-icon:ButtonText] border-transparent bg-[--btn-border] dark:bg-[--btn-bg] before:absolute before:inset-0 before:-z-10 before:rounded-[calc(theme(borderRadius.lg)-1px)] before:bg-[--btn-bg] before:shadow dark:before:hidden dark:border-white/5 after:absolute after:inset-0 after:-z-10 after:rounded-[calc(theme(borderRadius.lg)-1px)] after:shadow-[shadow:inset_0_1px_theme(colors.white/15%)] after:data-[active]:bg-[--btn-hover-overlay] after:data-[hover]:bg-[--btn-hover-overlay] dark:after:-inset-px dark:after:rounded-lg before:data-[disabled]:shadow-none after:data-[disabled]:shadow-none text-white [--btn-bg:theme(colors.zinc.900)] [--btn-border:theme(colors.zinc.950/90%)] [--btn-hover-overlay:theme(colors.white/10%)] dark:text-white dark:[--btn-bg:theme(colors.zinc.600)] dark:[--btn-hover-overlay:theme(colors.white/5%)] [--btn-icon:theme(colors.zinc.400)] data-[active]:[--btn-icon:theme(colors.zinc.300)] data-[hover]:[--btn-icon:theme(colors.zinc.300)] cursor-default`

four = `px-5 py-2.5 font-medium bg-blue-50 hover:bg-blue-100 hover:text-blue-600 text-blue-500 rounded-lg text-sm mt-8`
four = `

hover:bg-white
hover:py-16
rounded-full bg-black px-3.5 py-2.5 text-xs font-semibold text-white shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"`

setInterval(do_stuff, 1000)
function App() {


  const [count, setCount] = useState(0);
  const [isAccel, setIsAccel] = useState(true);

  const [imageSrc, setImageSrc] = useState('/most_recent.jpg');
  const [prevEvent, setPrevEvent] = useState({x: 0, y: 0})
  //let [backendUrl, setBackendUrl] = useState('')
    //  this.canvas = React.createRef()                  //// 1 - create ref
  const canvasRef = React.useRef(null);
  let backendUrl = ''

  useEffect(() => {
    const interval = setInterval(() => {
      // Append a timestamp to the image source to force re-fetching
      setImageSrc(`/most_recent.jpg?${new Date().getTime()}`);
      if (Math.random() > .9999999999) setCount(0)
    }, 150); // Refresh every 5000 milliseconds (5 seconds)

    return () => clearInterval(interval)     ; // Cleanup the interval on component unmount
  }, []);

  const handleMove = (shit: any) => {
    //console.log('move', shit)
    setPrevEvent(shit)

  }
  const handleStop = () => {
    //do_things()
    //speak()
  }

  const handleBrake = () => {
    send_brake_message()
  }



  const [isRunning, setIsRunning] = useState(false);
  const [intervalId, setIntervalId] = useState(null);

  const startRepeating = () => {
    if (!isRunning) {
      setIsRunning(true);
      const id = setInterval(() => {
        // Your function to call repeatedly
        console.log('Function is called repeatedly every second');
        //do_stuff()
        setShit(getShit + 1)
      }, 1000); // Change 1000ms (1 second) to your desired interval
      setIntervalId(id);
    }
  };

  const stopRepeating = () => {
    if (isRunning) {
      clearInterval(intervalId);
      setIsRunning(false);
      setIntervalId(null);
    }
  };
  const [getShit, setShit] = useState(1)
  // Optional: Cleanup on component unmount
  useEffect(() => {
    return () => intervalId && clearInterval(intervalId);
  }, [intervalId]);


  const base_name = `https://adnan-desktop-1.jerboa-kokanue.ts.net/api`

  const handleClick = () => {
    console.log('cick you me :)')
    fetch('/api/redeploy-webhook')


  }


  const handleArm = () => {
    console.log('cick you me :)')
    fetch(`${base_name}/redeploy-webhook`)
  }

  return (<div className="text-white">
    <h1 className="text-7xl">Robotics Teleoperation</h1>
    <div>1. roomba</div>
    <div>2. arm (trossen - order moar) </div>
    <div>3. zed camera</div>
    <div>4. jetson container vision transformer</div>
    <div>5. map, </div>
    {/* <div>nix setup?</div>
        <div>vm setup?</div>

        <div>cloud setup(fly.io)</div>
        <div>android atak setup</div>
        <div>map</div>
        <div>waypoint</div>
        <div>make 100 devices test.</div>
        <div>omni cam - scene understanding</div> */}
    {/* <img src={most_recent.jpg} height="500px" width="500px"/> */}

    <button className="bg-purple-500 mr-48" onClick={handleClick}>make shit drive</button>

    <button className="bg-purple-500" onClick={handleArm}>grasp arm</button>




  </div>);
  // <video src="https://adnan-dell-g16-7630.jerboa-kokanue.ts.net/vid.mp4" controls></video>
  return (
    <>
      <div className="">
        <div className={("")}>
          <div className={("" )}>
            <div className="w-full absolute">
              <div className="text-white absolute top-0 left-48">
                <div>x = {getX(prevEvent)}</div>
                <div>y = {getY(prevEvent)}</div>
                <div>Z = {getShit}</div>
                {/* <div>isAccel = {isAccel ? 'on': 'off'} </div>
                    <div>
                    {backendUrl ? backendUrl : ''}
                    </div> */}
              </div>
              <div className="">
                <iframe
                  width={window.innerWidth * .9}
                  className="m-auto aspect-video  hover:aspect-square" src="https://adnan-thinkpad-x1-carbon-gen-11.jerboa-kokanue.ts.net/"
                />
              </div>
              <div className={'absolute top-8 right-8 justify-between'} >
                <button className={four}

                        onMouseDown={() => startRepeating()} onMouseUp={stopRepeating}>Accel</button>
                <button className={four + ' ml-16'} onClick={handleBrake}>Brake</button>
              </div>
              <span className="absolute top-0 left-8">
              </span>
            </div>
            {/* <img
                style={{'margin': '0 auto'}}
                width="100px" height="100px"
                src="https://raw.githubusercontent.com/MarcoFazioRandom/Virtual-Joystick-Godot/Main/addons/virtual_joystick/previews/icon.png"
                /> */}
          </div>
        </div>
        {(count > 2) && (<Parking />) }
      </div>
    </>
  )
}

///<video controlvideoElements width="250"></video>
//_onClick={startRepeating}
export default App

// <button className={four} onClick={() => setBackendUrl(prompt('what server would you want - example https://5ac8-72-195-220-81.ngrok-free.app/drive') || '')}>server</button>
//ffmpeg -i rtsp://your_rtsp_stream_url -c copy -t 10 test.mp4


/* <button className={four}>Turn Left</button>
   <button className={four}>Turn Right</button>
   <button className={four}>Waypoint mode</button>
   <button
   onClick={() =>     do_things()          }
   className={four}>Session Replay</button> */
