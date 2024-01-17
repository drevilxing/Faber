import sys
from PyQt5.QtWidgets import QApplication, QMainWindow, QGraphicsScene, QGraphicsView, QGraphicsPixmapItem, QVBoxLayout, QWidget, QHBoxLayout, QPushButton, QLabel, QDialog, QLineEdit, QFormLayout
from PyQt5.QtCore import Qt, QMimeData, QPointF
from PyQt5.QtGui import QDrag, QPixmap
groups = []
class ImageDialog(QDialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setWindowTitle("Image Information")
        self.layout = QFormLayout(self)
        self.key = QLineEdit(self)
        self.blockchains = QLineEdit(self)
        self.channel = QLineEdit(self)
        self.layout.addRow("Key:", self.key)
        self.layout.addRow("Blockchains:", self.blockchains)
        self.layout.addRow("Channel:", self.channel)
        self.acceptButton = QPushButton("OK", self)
        self.acceptButton.clicked.connect(self.accept)
        self.layout.addRow(self.acceptButton)

    def get_data(self):
        return {
            "key": self.key.text(),
            "blockchains": self.blockchains.text(),
            "channel": self.channel.text()
        }

class GroupItem(QGraphicsPixmapItem):
    def __init__(self, pixmap, parent=None):
        super().__init__(pixmap, parent)
        self.group_data = {}

    def set_data(self, data):
        self.group_data = data

class ClickableGraphicsView(QGraphicsView):
    def __init__(self, scene, parent=None):
        super().__init__(scene, parent)
        self.groups = []

    def mousePressEvent(self, event):
        print("Mouse Pressed")  # 调试输出
        if event.button() == Qt.LeftButton:
            print("Left Button Clicked")  # 调试输出
            scene_pos = self.mapToScene(event.pos())
            pixmap = QPixmap("1.png")  # 确保图片路径正确
            item = GroupItem(pixmap)
            item.setPos(scene_pos - QPointF(pixmap.width() / 2, pixmap.height() / 2))
            self.scene().addItem(item)

            dialog = ImageDialog()
            result = dialog.exec_()  # 保存对话框返回值进行调试
            print("Dialog closed, result:", result)  # 调试输出
            if result:
                data = dialog.get_data()
                item.set_data(data)
                self.groups.append(item)


class DraggableLabel(QLabel):
    def __init__(self, shape, parent=None):
        super().__init__(parent)
        self.shape = shape
        self.initUI()

    def initUI(self):
        if self.shape == 'circle':
            self.setText("ca")
        elif self.shape == 'triangle':
            self.setText("orderer")
        elif self.shape == 'rectangle':
            self.setText("peer")
        self.setStyleSheet("border: 1px solid black;")

    def mousePressEvent(self, event):
        if event.button() == Qt.LeftButton:
            drag = QDrag(self)
            mimeData = QMimeData()

            mimeData.setText(self.shape)
            drag.setMimeData(mimeData)
            drag.exec_(Qt.MoveAction)

class ClickableGraphicsView(QGraphicsView):
    def __init__(self, scene, parent=None):
        super().__init__(scene, parent)

    def mousePressEvent(self, event):
        if event.button() == Qt.LeftButton:
            scene_pos = self.mapToScene(event.pos())
            pixmap = QPixmap("1.png")  # 确保此路径是正确的
            item = QGraphicsPixmapItem(pixmap)
            item.setPos(scene_pos - QPointF(pixmap.width()/2, pixmap.height()/2))  # 以图片中心为点击位置
            self.scene().addItem(item)


    

class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.initUI()

    def initUI(self):
        self.setWindowTitle("Blockchain Configuration Tool")
        self.setGeometry(100, 100, 800, 600)

        # Main layout
        hbox = QHBoxLayout()

        # Right side (Graphics view)
        self.scene = QGraphicsScene()
        self.view = ClickableGraphicsView(self.scene)
        self.view.setFixedSize(600, 600)
        hbox.addWidget(self.view)
# Left side (Toolbox)
        vbox = QVBoxLayout()

        # Draggable shapes
        self.circleLabel = DraggableLabel('circle')
        self.triangleLabel = DraggableLabel('triangle')
        self.rectangleLabel = DraggableLabel('rectangle')

        vbox.addWidget(self.circleLabel)
        vbox.addWidget(self.triangleLabel)
        vbox.addWidget(self.rectangleLabel)

        self.button = QPushButton("Generate", self)
        vbox.addWidget(self.button)
        hbox.addLayout(vbox)

        # Set main layout
        container = QWidget()
        container.setLayout(hbox)
        self.setCentralWidget(container)

# Run the application
if __name__ == '__main__':
    app = QApplication(sys.argv)
    mainWin = MainWindow()
    mainWin.show()
    sys.exit(app.exec_())

