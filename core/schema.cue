package core

#Template: {
   template: string @go(Template)
   path1: string @go(Path)
   path2: string @go(Path2)
}

#Config: {
   templates: [...#Template] @go(Templates,[]Template)
}

#Data: [string]: #Config
