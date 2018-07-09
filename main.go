package main

import(
//  "fmt"
  "log"
  "image"
  "image/color"
  "github.com/disintegration/imaging"
)

func main(){
   src, err := imaging.Open("flower.jpg")
   if err != nil{
     log.Fatal("Error %v", err)
   }

   //crop the image to square with center position and save it
   src = imaging.CropAnchor(src,300,300,imaging.Center)
   err = imaging.Save(src,"flowercrop.jpg")
   if err != nil{
     log.Fatal("Error %v",err)
   }

   //resize the cropped image and keep aspect ratio and save it
   src = imaging.Resize(src,200,0,imaging.Lanczos)
   err = imaging.Save(src,"resize.jpg")
   if err != nil{
     log.Fatal("Error %v",err)
   }

   //create four temp images and paste it on one canvas and save it

   //blur version of image
   img1 := imaging.Blur(src,5)

   //create grayscale version set higher contrass and sharpness to it.
   img2 := imaging.Grayscale(src)
   img2 = imaging.AdjustContrast(img2,20)
   img2 = imaging.Sharpen(img2,2)

   //invert image
   img3 := imaging.Invert(src)

   //color embossed version of pic
   img4 := imaging.Convolve3x3(src,
                              [9]float64{
                                        -1, -1, 0,
                                        -1, 1, 1,
                                         0, 1, 1,
                                       },nil)
  //create new canvas and save past four images on that
  canvas := imaging.New(400,400,color.NRGBA{0,0,0,0})
  canvas = imaging.Paste(canvas,img1,image.Pt(0,0))
  canvas = imaging.Paste(canvas,img2,image.Pt(0,200))
  canvas = imaging.Paste(canvas,img3,image.Pt(200,0))
  canvas = imaging.Paste(canvas,img4,image.Pt(200,200))

  err = imaging.Save(canvas,"canvas.jpg")
  if err != nil {
    log.Fatal("Error %v",err)
  }



}
