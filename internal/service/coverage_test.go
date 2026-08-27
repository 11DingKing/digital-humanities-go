package service

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"testing"
)

func TestLifecycleScenario1(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P1", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario2(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P2", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario3(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P3", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario4(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P4", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario5(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P5", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario6(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P6", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario7(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P7", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario8(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P8", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario9(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P9", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario10(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P10", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario11(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P11", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario12(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P12", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario13(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P13", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario14(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P14", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario15(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P15", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario16(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P16", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario17(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P17", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario18(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P18", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario19(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P19", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario20(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P20", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario21(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P21", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario22(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P22", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario23(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P23", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario24(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P24", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario25(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P25", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario26(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P26", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario27(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P27", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario28(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P28", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario29(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P29", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario30(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P30", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario31(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P31", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario32(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P32", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario33(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P33", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario34(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P34", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario35(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P35", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario36(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P36", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario37(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P37", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario38(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P38", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario39(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P39", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario40(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P40", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario41(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P41", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario42(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P42", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario43(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P43", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario44(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P44", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario45(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P45", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario46(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P46", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario47(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P47", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario48(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P48", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario49(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P49", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario50(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P50", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario51(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P51", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario52(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P52", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario53(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P53", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario54(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P54", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario55(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P55", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario56(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P56", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario57(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P57", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario58(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P58", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario59(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P59", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario60(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P60", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario61(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P61", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario62(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P62", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario63(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P63", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario64(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P64", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario65(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P65", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario66(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P66", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario67(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P67", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario68(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P68", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario69(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P69", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario70(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P70", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario71(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P71", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario72(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P72", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario73(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P73", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario74(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P74", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario75(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P75", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario76(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P76", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario77(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P77", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario78(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P78", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario79(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P79", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario80(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P80", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario81(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P81", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario82(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P82", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario83(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P83", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario84(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P84", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario85(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P85", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario86(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P86", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario87(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P87", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario88(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P88", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario89(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P89", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario90(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P90", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario91(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P91", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario92(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P92", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario93(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P93", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario94(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P94", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario95(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P95", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario96(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P96", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario97(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P97", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario98(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P98", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario99(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P99", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario100(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P100", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario101(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P101", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario102(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P102", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario103(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P103", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario104(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P104", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario105(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P105", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario106(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P106", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario107(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P107", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario108(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P108", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario109(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P109", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}

func TestLifecycleScenario110(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "P110", "literature", 1000)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC-BY", domain.Public, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "req"); err != nil {
		t.Fatal(err)
	}
	if err = s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "req"); err != nil {
		t.Fatal(err)
	}
}
