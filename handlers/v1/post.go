package v1

// PostIndex ...
// func PostIndex(c *fiber.Ctx) error {
// 	var currentPage = 1
// 	getCurrentPage := c.Query("page")
// 	if getCurrentPage != "" {
// 		currentPage, _ = strconv.Atoi(getCurrentPage)
// 	}
// 	limit := 15
// 	repo := repositories.NewPostRepository()
// 	r, err := repo.Paginate(currentPage, limit)

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": true,
// 		"data":   r,
// 	})
// }

// // PostCreate
// func PostCreate(c *fiber.Ctx) error {
// 	params := models.PostCreateInput{}

// 	c.BodyParser(&params)

// 	repo := repositories.NewPostRepository()
// 	r, err := repo.Create("a", params)

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"data": r,
// 	})
// }

// // PostShow
// func PostShow(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	repo := repositories.NewPostRepository()
// 	r, err := repo.FindById(id)

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": true,
// 		"data":   r,
// 	})
// }

// // PostUpdate
// func PostUpdate(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	params := models.PostUpdateInput{}

// 	c.BodyParser(&params)

// 	repo := repositories.NewPostRepository()
// 	r, err := repo.Update(id, params)

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": true,
// 		"data":   r,
// 	})
// }

// // PostDelete
// func PostDelete(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	repo := repositories.NewPostRepository()
// 	_, err := repo.Delete(id)

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status":  true,
// 		"message": utils.MsgSuccessDelete,
// 	})
// }
