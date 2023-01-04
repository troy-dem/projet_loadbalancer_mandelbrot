from flask import Flask, render_template, request
import aiohttp
import asyncio
import PIL.Image as Image
import io
import base64

app = Flask(__name__)
image_list = []
async def asyncCall(max_iteration,start_position_x,start_position_y,quantize_length):
    for n in range(100):
        image_list.append(" ")
    print(len(image_list))
    step = quantize_length/10.0
    async with aiohttp.ClientSession() as session:
        for i in range(10):
            for j in range(10):
                start_position_x_local = start_position_x + step*i
                start_position_y_local = start_position_y - step*j
                quantize_length_local = step
                params = {
                    "max_iteration" : max_iteration,
                    "start_position_x" : start_position_x_local,
                    "start_position_y" : start_position_y_local,
                    "quantize_length" : quantize_length_local
                }
                url_local = "http://localhost"
                async with session.get(url_local, params=params) as resp:
                    b64 = await resp.json()
                    b = base64.b64decode(b64["image"])
                    image_list[j*10+i] = b

@app.route('/')
def hello():
    return render_template("index.html")

@app.route('/mandelbrot')
def mandelbrot():
    im = Image.new('RGB',(500,500))
    im.save("static/images_loaded/mandelbrot_wanted.png")
    max_iteration = float(request.args.get('max_iteration'))
    start_position_x = float(request.args.get('start_position_x'))
    start_position_y  = float(request.args.get('start_position_y'))
    quantize_length = float(request.args.get('quantize_length'))
    asyncio.run(asyncCall(max_iteration,start_position_x,start_position_y,quantize_length))
    for number in range(100):
        img = Image.open(io.BytesIO(image_list[number]))
        base_image = Image.open("static/images_loaded/mandelbrot_wanted.png")
        base_image.paste(img,((number%10) * 50,(number//10)*50))
        base_image.save("static/images_loaded/mandelbrot_wanted.png")
    return render_template("index.html")
    # return str(quantize_length)

if __name__ == '__main__':
    app.run()

