package main 

import (
	"flag"
	"fmt"
	"os"
	"log"
	// "path/filepath"
)

func main() {
	
	textPtr := flag.String("c", "", "Type a controller name")
	createProj := flag.String("p", "", "Text to parse")
	flag.Parse()

	if *textPtr == "" && *createProj == ""  {
        flag.PrintDefaults()
        os.Exit(1)
    } else if *textPtr != "" {
		filecontext := "<?php\nnamespace " + *textPtr + "\\Controller;\n use Application\\Mvc\\Controller;\nclass " +*textPtr + "Controller extends Controller\n{\n//methods goes here...\n}\n"
		filename:= *textPtr + "Controller.php"
		CreateFile(filename, filecontext)
		os.Exit(1)
	} else if *createProj != ""{

		CreateDirIfNotExist(*createProj)
		CreateDirIfNotExist(*createProj + "/Controller")
		CreateDirIfNotExist(*createProj + "/Model")
		CreateDirIfNotExist(*createProj + "/Views/Admin")
		CreateDirIfNotExist(*createProj + "/Views/Index")
		CreateDirIfNotExist(*createProj + "/Widget")
		content := `<?php

		namespace ` +  *createProj + `\Controller;

		use Application\Mvc\Controller;
		use ` + *createProj + `\Model\` + *createProj + `;
		use ` + *createProj + `\Form\` + *createProj + `Form;
		
		class AdminController extends Controller
{

    public function initialize()
    {
        $this->setAdminEnvironment();
        $this->helper->activeMenu()->setActive('admin-portfolio');
    }

    public function indexAction()
    {

    }

    public function addAction()
    {

		{
			$this->view->pick(['admin/edit']);
			$form = new ` + *createProj + `Form();
			$model = new ` + *createProj + `();
			if ($this->request->isPost()) {
				$post = $this->request->getPost();
				$form->bind($post, $model);
	
				if ($form->isValid()) {
					if ($model->create()) {
						$form->bind($post, $model);
						if ($model->update()) {
							$this->uploadImage($model);
							$this->flash->success($this->helper->at('`+ *createProj + ` created'));
							return $this->redirect($this->url->get() . '` + *createProj + `/admin/edit/' . $model->getId() . '?lang=' . LANG);
							} else {
								$this->flashErrors($model);
							}
						} else {
							$this->flashErrors($model);
						}
					} else {
						$this->flashErrors($form);
					}
				}
		
				$this->view->model = $model;
				$this->view->form = $form;
		
				$this->helper->title($this->helper->at('Create a ` + *createProj +`'), true);

			}
		
			public function editAction($id)
			{
				$id = (int) $id;
				$form = new ` + *createProj + `Form();
				$model = ` + *createProj + `::findFirst($id);

				if ($this->request->isPost()) {
					$post = $this->request->getPost();
					$form->bind($post, $model);
					if ($form->isValid()) {
						if ($model->save()) {
							$this->flash->success($this->helper->at('` + *createProj + ` edited'));

							return $this->redirect($this->url->get() . '` + *createProj + `/admin/edit/' . $model->getId() . '?lang=' . LANG);
							} else {
								$this->flashErrors($model);
							}
						} else {
							$this->flashErrors($form);
						}
					} else {
						$form->setEntity($model);
					}
			
					$this->view->model = $model;
					$this->view->form = $form;
					$this->helper->title($this->helper->at('Edit ` + *createProj + `'), true);
				}
			
				public function deleteAction($id)
				{
			
				}
			}
		}`
		CreateFile(*createProj + "/Controller/AdminController.php", content)
		ProjNameWidgetContent := `<?php

namespace ` + *createProj + `\Widget;

use Application\Widget\AbstractWidget;
use ` + *createProj + `\Model\` + *createProj + `;


class ` + *createProj + `Widget extends AbstractWidget
{

    public function portfolio($limit = 6)
    {
    	$parameters['order'] = "sort ASC";
        $portfolio = ` + *createProj + `::find([
			'limit' => ['number' => $limit]
		  ]);
		  $this->widgetPartial('index/block', ['portfolio' => $portfolio]);
	  }
  
  }
  `
  CreateFile(*createProj + "/Widget/" +*createProj + "Widget.php", ProjNameWidgetContent)

  ProjNameModelContent := `<?php

  namespace ` + *createProj + `\Widget;

  use Application\Widget\AbstractWidget;
  use ` + *createProj + `\Model;

  use Application\Cache\Keys;
  use Application\Mvc\Model\Model;
  use Phalcon\Validation;
  use Phalcon\Validation\Validator\Uniqueness as UniquenessValidator;
  use Application\Localization\Transliterator;
  
  class ` + *createProj + `extends Model
  {
  
	  public function getSource()
	  {
		  //Data base table name
		  return "` + *createProj + `";
		}
	
		public function initialize()
		{
	
		}
	
		private $id;
	
		public function getId()
		{
			return $this->id;
		}
	}`
	CreateFile(*createProj + `\Model\` + *createProj + `.php`, ProjNameModelContent)
	}
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
					panic(err)
			}
	}
}

func CreateFile(name, content string){
	file, err := os.Create(name)
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer file.Close()
		fmt.Fprintf(file, content)
}