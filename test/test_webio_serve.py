import tornado.ioloop
import tornado.web
from pywebio.input import input, FLOAT
from pywebio.output import put_text
from pywebio.platform.tornado import webio_handler


class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, world")


def bmi():
    height = input("Input your height(cm)：", type=FLOAT)
    weight = input("Input your weight(kg)：", type=FLOAT)

    BMI = weight / (height / 100) ** 2

    top_status = [(16, 'Severely underweight'), (18.5, 'Underweight'),
                  (25, 'Normal'), (30, 'Overweight'),
                  (35, 'Moderately obese'), (float('inf'), 'Severely obese')]

    for top, status in top_status:
        if BMI <= top:
            put_text('Your BMI: %.1f. Category: %s' % (BMI, status))
            break


if __name__ == "__main__":
    application = tornado.web.Application([
        (r"/", MainHandler),
        (r"/bmi", webio_handler(bmi)),  # bmi is the same function as above
    ])
    application.listen(port=80, address='localhost')
    tornado.ioloop.IOLoop.current().start()
