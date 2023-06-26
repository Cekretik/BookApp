package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gin/gin-gonic"
	"strconv"
	"github.com/Cekretik/BookApp/main/pkg/utils"
	"github.com/Cekretik/BookApp/main/pkg/models"
)

var NewBook models.Bo